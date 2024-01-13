import { redirect } from "@sveltejs/kit"


export const load = async ({ locals, params, cookies }) => {
    let user_not_signin = true

    let jwt = cookies.get("jwt_lab_rank")
    let collegeID = locals.user.college_id
    let subjectID = params.slug
    let userID = locals.user.id
    let languages = []
    if (jwt != undefined) {
        user_not_signin = false
    }

    let environments = []
    const listEnvironments = async () => {
        try {
            console.log("server listEnvironments")
            const response = await fetch(`http://127.0.0.1:8080/environment`)
            const data = await response.json();
            environments = data.Message;
            // console.log("sfsfa", environments);
        } catch (error) {
            console.error("Error fetching environments:", error);
        }
    };
    let syllabus = []
    const listSyllabus = async () => {
        try {
            console.log("server list syllabus")
            const response = await fetch(`http://127.0.0.1:8080/syllabus/by_subject/${subjectID}`)
            const data = await response.json();
            syllabus = data.Message;
            // console.log("sfsfa", syllabus);
        } catch (error) {
            console.error("Error fetching syllabus:", error);
        }
    };
    console.log(user_not_signin)
    await listEnvironments()
    await listSyllabus()

    const environmentMap = environments.map((env) => ({ id: env.id, title: env.title }));
    const syllabusMap = syllabus.map((slb) => ({ id: slb.id, syllabus_level: slb.syllabus_level }));

    console.log(environmentMap)
    return {
        user_not_signin,
        collegeID,
        subjectID,
        environmentMap,
        userID,
        syllabusMap
    };

}


export const actions = {
    create: async ({ fetch, locals, request }) => {
        const data = await request.formData();
        console.log(data)
        let languages = [];
        const environments = data.getAll('environments').map(envString => {
            const [id, language] = envString.split('_');
            languages.push(language)
            return {
                id,
                language,
            };
        });
        languages = [...new Set(languages)];
        console.log(languages)
        const testFiles = [];

        data.forEach((value, key) => {
            const match = key.match(/^testFile(?:Title|InitCode)?_(\w+)$/);
            // console.log(match)
            // console.log(match[1])
            if (match && languages.includes(match[1])) {
                var index = languages.indexOf(match[1])
                // console.log(key, match[1], match[1] in languages, languages)
                const language = match[1];

                if (!testFiles[index]) {
                    testFiles[index] = {
                        language: language,
                        title: '',
                        init_code: '',
                        file: '',
                    };
                }

                if (key.startsWith(`testFile_${language}`)) {
                    testFiles[index].file = btoa(value);
                } else if (key.startsWith(`testFileInitCode_${language}`)) {
                    testFiles[index].init_code = btoa(value);
                } else if (key.startsWith(`testFileTitle_${language}`)) {
                    testFiles[index].title = value;
                }
            }
        });

        console.log(testFiles)

        // Remove any undefined elements from the array
        const filteredTestFiles = testFiles.filter(Boolean);


        const jsonData = {
            title: data.get('title'),
            created_by: locals.user.id,
            difficulty: data.get('difficulty'),
            syllabus_id: data.get('syllabusId'),
            environment:
                environments,
            problem_file: btoa(data.get('problemFile')),
            test_files: testFiles
        };

        console.log(jsonData)

        const response = await fetch('http://127.0.0.1:8080/problem', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(jsonData),
        });

        // console.log(response);
        if (response.ok) {
            throw redirect(300, "/subjects")
        }

    }

};


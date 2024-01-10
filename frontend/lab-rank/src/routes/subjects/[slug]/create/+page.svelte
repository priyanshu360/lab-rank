<script>
  import { onMount } from "svelte";
  import Header from "$lib/Header.svelte";
  import Footer from "$lib/Footer.svelte";
  import Description from "$lib/Description.svelte";
  let problemData = {
    title: "",
    difficulty: "",
    syllabusId: "",
    environments: [], // Use an array to store multiple environments
    testFiles: [],
    file: "",
  };

  export let data;

  // Dropdown options
  let environmentOptions = data.environmentMap;

  let syllabusOptions = data.syllabusMap;

  onMount(() => {
    // Initialize the component with test data
    initializeTestFiles();
  });

  const initializeTestFiles = () => {
    console.log(problemData.testFiles);
    environmentOptions.forEach((env) => {
      problemData.testFiles.push({
        language: env.title,
        title: "",
        init_code: "",
        file: "",
      });
      console.log(problemData.testFiles);
    });
  };

  console.log(data.environmentMap, environmentOptions);
  // Function to update the test file language based on the selected environment
  const updateTestFileLanguage = (language) => {
    problemData.testFiles[0].language = language;
  };

  function handleFileChange(e) {
    const file = e.target.files[0];
    console.log(" handle file change 2");
    if (file) {
      const reader = new FileReader();
      reader.onload = function (event) {
        problemData.file = event.target.result;
      };
      reader.readAsText(file);
    } else {
      problemData.file = "";
    }
  }

  function handleFileChangeAtIndex(e, index) {
    const file = e.target.files[0];
    console.log(e);
    if (file) {
      const reader = new FileReader();
      reader.onload = function (event) {
        problemData.testFiles[index].file = event.target.result;
        console.log(
          "File content for index",
          index,
          ":",
          problemData.testFiles[index].file
        );
      };
      reader.readAsText(file);
      console.log("affsf", problemData.testFiles);
    } else {
      problemData.testFiles[index].file = "";
    }
  }
</script>

<Header {data} />
<Description />
<main>
  <h1>Create Problem</h1>

  <form action="?/create" method="POST">
    <!-- Other form fields -->
    <label>
      Title:
      <input type="text" bind:value={problemData.title} name="title" />
    </label>

    <label>
      Difficulty:
      <select bind:value={problemData.difficulty} name="difficulty">
        <option value="">Select Difficulty</option>
        <option value="EASY">Easy</option>
        <option value="MEDIUM">Medium</option>
        <option value="HARD">Hard</option>
      </select>
    </label>

    <label>
      Syllabus Level:
      <select bind:value={problemData.syllabusId} name="syllabusId">
        {#each syllabusOptions as slb}
          <option value={slb.id}>{slb.syllabus_level}</option>
        {/each}
      </select>
    </label>

    <label>
      Environments:
      {#each environmentOptions as env}
        <label>
          <input
            type="checkbox"
            bind:group={problemData.environments}
            value={env.id + "_" + env.title}
            name="environments"
          />
          {env.title}
        </label>
      {/each}
    </label>

    <label>
      Problem File:
      <input type="file" accept=".txt, .json" on:change={handleFileChange} />
    </label>

    <label>
      <input hidden name="problemFile" bind:value={problemData.file} />
    </label>
    <label>
      Test Script:
      {#each problemData.testFiles as { language, title, init_code, file, ...rest }, index (language)}
        <div>
          <label>
            Test File for {language}:
            <input
              type="file"
              accept=".txt, .json"
              on:change={(e) => handleFileChangeAtIndex(e, index)}
            />
          </label><br />

          <label>
            Title:
            <input
              type="text"
              hidden
              value={`${language}_${problemData.title}`}
              name={`testFileTitle_${language}`}
            />
          </label><br />
          <label>
            Init Code:
            <textarea
              rows="5"
              cols="33"
              bind:value={problemData.testFiles[index].init_code}
              name={`testFileInitCode_${language}`}
            ></textarea>
          </label><br />
        </div>
        <label>
          <input
            name={`testFile_${language}`}
            hidden
            bind:value={problemData.testFiles[index].file}
          /></label
        >
        <br />
      {/each}
    </label>

    <button type="submit">Create Problem</button>
  </form>
</main>
<Footer />

<style>
  /* Add your styles here */
  main {
    /* text-align: center; */
    max-width: 500px;
    margin: 0 auto;
    padding: 1rem;
    font-family: "Khand", sans-serif;
    background-color: #f4f4f4;
  }
  button {
    background-color: #04aa6d;
    margin-bottom: 20px;
  }
  form {
    display: grid;
    gap: 1.5rem;
  }
</style>

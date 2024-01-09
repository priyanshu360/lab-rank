<script>
  import { onMount } from "svelte";
  import Header from "$lib/Header.svelte";
  import Footer from "$lib/Footer.svelte";
  import Description from "$lib/Description.svelte";
  let problemData = {
    title: "",
    createdBy: "",
    difficulty: "",
    syllabusId: "",
    environments: [], // Use an array to store multiple environments
    testFiles: [],
  };

  export let data;
  // Example data for testing
  let testData = {
    title: "TEST",
    createdBy: "2ddc8a9e-c558-4246-ae5d-1964c1dedf62",
    difficulty: "MEDIUM",
    syllabusId: "a51a659e-ec11-4d84-ae67-39b3c3e70820",
    environments: [
      {
        language: "Go",
        id: "9809a649-9200-4e85-aaab-12bab28947c8",
      },
      {
        language: "JavaScript",
        id: "1234",
      },
      // Add more environments as needed
    ],
    testFiles: [],
  };

  // Dropdown options
  let environmentOptions = [
    { id: "1", language: "JavaScript" },
    { id: "2", language: "Python" },
    { id: "3", language: "Java" },
    { id: "4", language: "Go" },
    { id: "5", language: "C++" },
  ];

  let syllabusOptions = [
    { id: "a51a659e-ec11-4d84-ae67-39b3c3e70820", name: "Syllabus 1" },
    { id: "another-syllabus-id", name: "Syllabus 2" },
    // Add more syllabus options as needed
  ];

  onMount(() => {
    // Initialize the component with test data
    problemData = testData;
    initializeTestFiles();
  });

  const handleSubmit = () => {
    // Implement your logic to handle form submission
    console.log("Submitting problem data:", problemData);
    // Add your API call or further processing logic here
  };

  const initializeTestFiles = () => {
    console.log(problemData.testFiles);
    environmentOptions.forEach((env) => {
      problemData.testFiles.push({
        language: env.language,
        title: "",
        init_code: "",
        file: "",
      });
      console.log(problemData.testFiles);
    });
  };

  initializeTestFiles;
  // Function to update the test file language based on the selected environment
  const updateTestFileLanguage = (language) => {
    problemData.testFiles[0].language = language;
  };
</script>

<Header {data} />
<Description />
<main>
  <h1>Create Problem</h1>

  <form on:submit|preventDefault={handleSubmit}>
    <!-- Other form fields -->
    <label>
      Title:
      <input bind:value={problemData.title} />
    </label>

    <label>
      Created By:
      <input bind:value={problemData.createdBy} />
    </label>

    <label>
      Difficulty:
      <input bind:value={problemData.difficulty} />
    </label>

    <label>
      Syllabus ID:
      <select bind:value={problemData.syllabusId}>
        {#each syllabusOptions as { id, name }}
          <option value={id}>{name}</option>
        {/each}
      </select>
    </label>

    <label>
      Environments:
      {#each environmentOptions as { id, language }}
        <label>
          <input
            type="checkbox"
            bind:group={problemData.environments}
            value={id}
          />
          {language}
        </label>
      {/each}
    </label>

    <label>
      Problem File:
      <input
        type="file"
        accept=".txt, .json"
        on:change={(e) => (problemData.testFiles[0].file = e.target.files[0])}
      />
    </label>

    <label>
      Test Script:
      {#each problemData.testFiles as { language, title, init_code, file, ...rest }, index (language)}
        <div bind:this={problemData.testFiles[index]}>
          <label>
            Test File for {language}:
            <input
              type="file"
              accept=".txt, .json"
              on:change={(e) => {
                problemData.testFiles[index].file = e.target.files[0];
                updateTestFileLanguage(index, language);
              }}
            />
          </label><br />
          <label>
            Title:
            <input type="text" bind:value={title} />
          </label><br />
          <label>
            Init Code:
            <textarea rows="5" cols="33" bind:value={init_code}></textarea>
          </label><br />
        </div>
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

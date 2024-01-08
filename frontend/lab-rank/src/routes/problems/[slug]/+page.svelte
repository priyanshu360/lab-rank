<script>
  import { onMount } from "svelte";
  import CodeMirror from "svelte-codemirror-editor";
  import { javascript } from "@codemirror/lang-javascript";
  import { python } from "@codemirror/lang-python";
  import { sql } from "@codemirror/lang-sql";
  import Header from "../../../lib/Header.svelte";
  import Footer from "../../../lib/Footer.svelte";
  import Description from "../../../lib/Description.svelte";
  // import { oneDark } from "@codemirror/theme-one-dark";

  export let data;
  export let problem_title = data.responseData.Message.title;
  export let problem_file = atob(data.responseData.Message.problem_file);

  let languages = data.responseData.Message.environment.map(
    (env) => env.language
  ); // Replace this with your array of language options
  let selectedLanguage = languages[0];

  let languageMap = {
    JavaScript: javascript(),
    Python: python(),
    MySql: sql(),
  };

  export let code;

  const getInitCode = async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/problem?id=${data.slug}&lang=${selectedLanguage}`
      );
      const responseData = await response.json();
      console.log(responseData);

      code = atob(responseData.Message);
    } catch (error) {
      console.error("Error fetching problem details:", error);
    }
  };

  const submitCode = async () => {
    try {
      const response = await fetch("http://localhost:8080/submission", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          problem_id: data.slug,
          created_by: data.userID,
          score: null,
          run_time: null,
          metadata: {},
          lang: "Go",
          status: "QUEUED",
          solution: btoa(code),
        }),
      });

      console.log(data.slug);

      const responseData = await response.json();
      console.log("Submission response:", responseData);
    } catch (error) {
      console.error("Error submitting code:", error);
    }
  };

  // This function is called when the CodeMirror editor instance is ready
  const editorReady = (editor) => {
    // Access the CodeMirror instance through the 'editor' variable
    // You can perform additional configuration or operations here
    // Example: editor.setOption('someOption', 'someValue');
  };
</script>

<Header {data} />
<Description />
<main>
  <section class="problem-description">
    <h2>{problem_title}</h2>
    <p>{@html problem_file}</p>
  </section>

  <form class="code-editor-container" action="?/create">
    <label for="language-select">Select Language:</label>
    <select
      id="language-select"
      bind:value={selectedLanguage}
      on:change={getInitCode}
      name="language"
    >
      {#each languages as lang (lang)}
        <option value={lang}>{lang}</option>
      {/each}
    </select>
    <CodeMirror
      name="code"
      styles={{
        "&": {
          // width: "500px",
          // maxWidth: "100%",
          height: "50rem",
        },
      }}
      bind:value={code}
      lang={languageMap[selectedLanguage]}
      options={{
        lineNumbers: true,
        theme: "default", // Adjust the theme based on your preferences
      }}
      on:editorReady={editorReady}
    />

    <button on:click={submitCode}>Submit Code</button>
  </form>
</main>

<Footer />

<style>
  main {
    display: flex;
    justify-content: space-between; /* Changed to space-between for better distribution */
    align-items: flex-start;
    margin: 20px;
    font-family: "Khand", sans-serif;
  }

  .problem-description {
    flex: 1;
    /* max-width: 800px; */
    max-height: 1000px;
    overflow: auto;
    word-wrap: break-word;
    flex-direction: column;
    padding-right: 20px; /* Added padding to the right for better spacing */
  }

  .code-editor-container {
    flex: 1;
    margin-top: 20px; /* Adjusted margin-top for better spacing */
    padding: 20px;
    /* max-width: 600px; */
    display: flex;
    flex-direction: column;
  }

  label {
    display: -webkit-inline-box;
    margin-bottom: 5px; /* Added margin-bottom for better spacing between label and select */
  }

  select {
    max-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
    padding: 12px 16px;
    z-index: 1;
    margin-bottom: 10px;
    display: -webkit-inline-box;
    text-align: center;
    text-decoration: none;
    display: inline-block;
  }

  button {
    max-width: 160px;
    margin-top: 10px;
    padding: 20px;
    cursor: pointer;
    background-color: #04aa6d; /* Green */
    border: none;
    color: white;
    text-align: center;
    text-decoration: none;
    display: flex;
    font-size: 16px;
  }
</style>

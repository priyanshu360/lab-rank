<script>
  import { onMount } from "svelte";
  import CodeMirror from "svelte-codemirror-editor";
  import { javascript } from "@codemirror/lang-javascript";
  import Header from "../../../lib/Header.svelte";
  import Footer from "../../../lib/Footer.svelte";
  import Description from "../../../lib/Description.svelte";

  export let data;
  export let problem_title;
  export let problem_file;

  let code = "";

  onMount(async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/problem?id=${data.slug}`
      );
      const responseData = await response.json();

      problem_title = responseData.Message.title;
      problem_file = atob(responseData.Message.problem_file);
      console.log(problem_file);
    } catch (error) {
      console.error("Error fetching problem details:", error);
    }
  });

  const submitCode = () => {
    // Access the code through the 'code' variable
    console.log("Code submitted:", code);
  };

  // This function is called when the CodeMirror editor instance is ready
  const editorReady = (editor) => {
    // Access the CodeMirror instance through the 'editor' variable
    // You can perform additional configuration or operations here
    // Example: editor.setOption('someOption', 'someValue');
  };
</script>

<Header />
<Description />

<main>
  <section class="problem-description">
    <h2>{problem_title}</h2>
    <p>{problem_file}</p>
  </section>

  <section class="code-editor-container">
    <CodeMirror
      bind:value={code}
      lang={javascript()}
      options={{
        lineNumbers: true,
        theme: "default", // Adjust the theme based on your preferences
      }}
      styles={{
        "&": {
          height: "20rem",
        },
      }}
      on:editorReady={editorReady}
    />

    <button on:click={submitCode}>Submit Code</button>
  </section>
</main>

<Footer />

<style>
  main {
    display: flex;
    justify-content: space-around;
    align-items: flex-start;
    margin: 20px;
    font-family: "Khand", sans-serif;
  }

  .problem-description {
    flex: 1;
    max-width: 600px;
    /* max-height: 1000px; Adjust the max-height based on your design */
    overflow: auto; /* Use auto or scroll based on your preference */
    word-wrap: break-word; /* Allow words to break and wrap onto the next line */
  }
  .code-editor-container {
    flex: 1;
    margin-top: 100px;
    max-width: 600px;
    margin-left: 20px;
    display: flex;
    flex-direction: column;
  }

  button {
    margin-top: 10px;
    padding: 10px;
    cursor: pointer;
  }
</style>

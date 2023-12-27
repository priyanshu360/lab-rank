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

  export let code;

  onMount(async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/problem?id=${data.slug}`
      );
      const responseData = await response.json();

      problem_title = responseData.Message.title;
      problem_file = atob(responseData.Message.problem_file);
    } catch (error) {
      console.error("Error fetching problem details:", error);
    }
    try {
      const response = await fetch(
        `http://localhost:8080/problem?id=${data.slug}&lang=Go`
      );
      const responseData = await response.json();
      console.log(responseData);

      code = atob(responseData.Message);
    } catch (error) {
      console.error("Error fetching problem details:", error);
    }
  });

  const submitCode = async () => {
    try {
      const response = await fetch("http://localhost:8080/submission", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          problem_id: data.slug,
          link: "submission_link",
          created_by: "f77ad338-7fe7-4093-bdbb-c24ec489e10c",
          score: null,
          run_time: null,
          metadata: {},
          lang: "Go",
          status: "QUEUED",
          solution: btoa(code),
        }),
      });

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

  <section class="code-editor-container">
    <CodeMirror
      bind:value={code}
      lang={javascript()}
      options={{
        lineNumbers: true,
        theme: "default", // Adjust the theme based on your preferences
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
    max-width: 800px;
    max-height: 1000px;
    overflow: auto;
    display: flex;
    word-wrap: break-word; /* Allow words to break and wrap onto the next line */
    flex-direction: column;
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

<!-- SubjectPage.svelte -->

<script>
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import SubjectList from "./SubjectList.svelte";
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";
  import CreateSubjectModal from "./CreateSubjectModal.svelte";

  import { sessionStore } from '../../lib/sessionStore.js';

  export let data;
  
  sessionStore.set({
    subjects: data.subjects || [], // Adding Fetched Subjects into sessionStore 
  });

  let sessionData;
  sessionStore.subscribe(value => sessionData = value);

  let subjects;
  $: subjects = sessionData.subjects;
  
  let universities = data.universities; 


  let showModal = false;

  if (browser && data.user_not_signin == true) {
    goto("/signup");
  }

  function openModal() {
    showModal = true;
  }

  function closeModal() {
    showModal = false;
    apiStatus = null;
  }

  async function handleSubmit(formData) {
    const apiEndpoint = 'http://localhost:8080/subject';
    const apiRequestBody = {
      title: formData.detail.subjectTitle,
      description: formData.detail.subjectDescription,
      university_id: formData.detail.universityID,
    };

    try {
      const apiResponse = await fetch(apiEndpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(apiRequestBody),
      });

      if (apiResponse.ok) {
        sessionStore.update(currentSessionData => ({
          ...currentSessionData,
          subjects: [...currentSessionData.subjects, apiRequestBody],
        }));
      } else {
        console.error('API call failed');
      }
    } catch (error) {
      console.error('Error making API call:', error);
    }
    closeModal();
  }

</script>

<!-- Remaining HTML and styling unchanged -->
<Header {data} />
<Description />
<div class="center">
  <button class="create-button" on:click={openModal}>Create Subject</button>
</div>
<SubjectList {subjects} />
<Footer />


{#if showModal}
  <CreateSubjectModal {universities} on:submit={handleSubmit} on:closeModal={closeModal} />
{/if}


<style>
  /* Center the button horizontally */
  .center {
    display: flex;
    justify-content: center;
    margin-top: 20px; /* Adjust the margin as needed */
  }

  /* Style the button to have a button-like appearance */
  .create-button {
    /* Add button-like styles */
    background-color: #007bff; /* Set your preferred button color */
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
    /* Add more styles as needed */
  }
</style>

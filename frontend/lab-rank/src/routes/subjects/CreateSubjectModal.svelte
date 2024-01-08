<!-- CreateSubjectModal.svelte -->

<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let universities = [] ;
  let subjectTitle = "";
  let subjectDescription = "";
  let universityID = "";

  function handleSubmit() {
    const formData = {
      subjectTitle,
      subjectDescription,
      universityID,
    };

    dispatch("submit", formData);
    closeModal();
  }

  function closeModal() {
    dispatch("closeModal");
  }
</script>

<div class="modal">
  <div class="modal-content">
    <h2>Create Subject</h2>
    <form on:submit|preventDefault={handleSubmit}>
      <label>
        Subject Title:
        <input type="text" bind:value={subjectTitle} />
      </label>
      <label>
        Subject Description:
        <textarea bind:value={subjectDescription}></textarea>
      </label>
      <label>
        University:
        <select bind:value={universityID} name="universityID">
          <option value="" disabled selected>Select your option</option>
          {#each universities as university}
            <option value={university.ID}>{university.Title}</option>
          {/each}
        </select>
      </label>
      <!-- Add more input fields for other subject details -->

      <button type="submit" on:submit={handleSubmit}>Submit</button>
      <button type="button" on:click={closeModal}>Close</button>
    </form>
  </div>
</div>

<style>
  /* Add your modal styles here */
  .modal {
    /* Styles for the modal container */
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .modal-content {
    /* Styles for the modal content */
    background-color: white;
    padding: 20px;
    border-radius: 5px;
  }

  /* Add more styles for input fields, labels, etc., as needed */
  label {
    display: block;
    margin-bottom: 10px;
  }

  input[type="text"],
  textarea {
    width: 100%;
    padding: 8px;
    margin-bottom: 10px;
    border-radius: 4px;
    border: 1px solid #ccc;
  }

  button {
    padding: 8px 16px;
    border-radius: 4px;
    border: none;
    cursor: pointer;
    margin-right: 10px;
  }
</style>

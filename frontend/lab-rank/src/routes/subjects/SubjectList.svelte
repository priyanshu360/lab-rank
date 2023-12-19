<!-- SubjectList.svelte -->

<script>
  export let subjects = [];

  // Optional: You can add more properties as needed
</script>

<main class="max-w-2xl mx-auto p-8">
  <h1 class="text-2xl font-bold mb-4">Select Subject</h1>

  {#if subjects.length === 0}
    <p>No subjects available.</p>
  {:else}
    <div class="grid gap-4">
      {#each subjects as { id, title, description, university_id }}
        <div class="bg-white p-4 rounded shadow-md">
          <h2 class="text-xl font-bold mb-2">{title}</h2>
          <p class="text-gray-600">{description}</p>
          <p class="text-gray-500 mt-2">University ID: {university_id}</p>
        </div>
      {/each}
    </div>
  {/if}
</main>

<style>
  main {
    text-align: center;
  }

  /* Add additional styling for cards as needed */
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 1rem;
  }

  .rounded {
    border-radius: 0.5rem;
  }

  .shadow-md {
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
</style>

<script>
  import Header from "$lib/Header.svelte";
  import Footer from "$lib/Footer.svelte";
  import Table from "$lib/Table.svelte";
  import Description from "$lib/Description.svelte";
  export let data;
  export let subjectID = data.subjectID;
  export let collegeID = data.collegeID;
</script>

<main>
  <!-- Your main content goes here -->

  <Header {data} />

  <!-- Rest of your content -->
  <Description />
  <Table {subjectID} {collegeID} />

  <Footer />
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }
</style>

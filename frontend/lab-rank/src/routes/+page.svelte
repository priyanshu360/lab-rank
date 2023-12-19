<script>
  import Header from "../lib/Header.svelte";
  import Footer from "../lib/Footer.svelte";
  import Table from "../lib/Table.svelte";
  import Description from "../lib/Description.svelte";
</script>

<main>
  <!-- Your main content goes here -->

  <Header />

  <!-- Rest of your content -->
  <Description />
  <Table />

  <Footer />
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }
</style>

<!-- Login.svelte -->

<script>
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";

  export let data;
</script>

<Header {data} />
<Description />
<main class="text-center max-w-2xl mx-auto p-8">
  <h1 class="text-2xl font-bold mb-4">Login</h1>
  <form action="?/create" method="POST" class="grid gap-4">
    <label>
      Email:
      <input type="email" name="email" />
    </label>

    <label>
      Password:
      <input type="password" name="password" />
    </label>

    <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded"
      >Login</button
    >
  </form>
</main>

<Footer />

<style>
  main {
    text-align: center;
    max-width: 500px;
    margin: 0 auto;
    font-family: "Khand", sans-serif;
    padding: 2rem;
  }

  form {
    display: grid;
    gap: 1rem;
  }
</style>

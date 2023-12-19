<!-- Login.svelte -->

<script>
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";
  let email = "";
  let password = "";

  const handleSubmit = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
        }),
      });

      const data = await response.json();

      // Handle the login response as needed
      console.log(data);
    } catch (error) {
      console.error("Error during login:", error);
    }
  };
</script>

<Header />
<Description />
<main class="text-center max-w-2xl mx-auto p-8">
  <h1 class="text-2xl font-bold mb-4">Login</h1>
  <form on:submit|preventDefault={handleSubmit} class="grid gap-4">
    <label>
      Email:
      <input bind:value={email} type="email" />
    </label>

    <label>
      Password:
      <input bind:value={password} type="password" />
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
    max-width: 240px;
    margin: 0 auto;
    font-family: "Khand", sans-serif;
    padding: 2rem;
  }

  form {
    display: grid;
    gap: 1rem;
  }
</style>

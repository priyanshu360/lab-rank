<script>
  import { onMount } from "svelte";
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";

  let college_id = "";
  let email = "";
  let contact_no = "";
  let dob = "";
  let university_id = "";
  let name = "";
  let user_name = "";
  let password = "";

  let universities = [
    { id: "fdf3a549-5ed7-40f4-8706-9da30265aa7a", name: "University A" },
    { id: "other-university-id", name: "University B" },
    // Add more universities as needed
  ];

  // const fetchUniversities = async () => {
  //   try {
  //     // const response = await fetch("http://localhost:8080/api/universities");
  //     // const data = await response.json();
  //     universities = [
  //       { id: "fdf3a549-5ed7-40f4-8706-9da30265aa7a", name: "University A" },
  //       { id: "other-university-id", name: "University B" },
  //       // Add more universities as needed
  //     ]; // Assuming the API returns an array of universities
  //   } catch (error) {
  //     console.error("Error fetching universities:", error);
  //   }
  // };

  const handleSubmit = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          college_id,
          email,
          contact_no,
          dob,
          university_id,
          name,
          user_name,
          password,
        }),
      });

      const data = await response.json();
      console.log(data); // Handle the response as needed
    } catch (error) {
      console.error("Error during signup:", error);
    }
  };

  // onMount(() => {
  //   // Fetch universities when the component mounts
  //   // fetchUniversities();
  // });
</script>

<!-- Remaining HTML and styling unchanged -->
<Header />
<Description />

<main class="text-center max-w-2xl mx-auto p-8">
  <h1 class="text-2xl font-bold mb-4">Create New Account</h1>
  <form on:submit|preventDefault={handleSubmit} class="grid gap-4">
    <label>
      College ID:
      <input bind:value={college_id} type="text" />
    </label>

    <label>
      Email:
      <input bind:value={email} type="email" />
    </label>

    <label>
      Contact Number:
      <input bind:value={contact_no} type="tel" />
    </label>

    <label>
      Date of Birth:
      <input bind:value={dob} type="date" />
    </label>

    <label>
      University:
      <select bind:value={university_id} class="border p-2 mt-1">
        {#each universities as { id, name }}
          <option value={id}>{name}</option>
        {/each}
      </select>
    </label>

    <label>
      Name:
      <input bind:value={name} type="text" />
    </label>

    <label>
      User Name:
      <input bind:value={user_name} type="text" />
    </label>

    <label>
      Password:
      <input bind:value={password} type="password" />
    </label>

    <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded"
      >Sign Up</button
    >
  </form>
</main>

<Footer />

<style>
  main {
    text-align: center;
    max-width: 240px;
    margin: 0 auto;
    padding: 2rem;
    font-family: "Khand", sans-serif;
  }

  form {
    display: grid;
    gap: 1rem;
  }
</style>

<script>
  import { onMount } from "svelte";
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";
  import { format } from "date-fns";

  let college_id = "";
  let email = "";
  let contact_no = "";
  let dob = "";
  let university_id = "";
  let name = "";
  let user_name = "";
  let password = "";
  let fdob = "";

  let colleges = [];
  let universities = [];

  let emailValid = true;
  let contactNoValid = true;

  const formattedDob = () => {
    dob = format(new Date(fdob), "yyyy-MM-dd'T'HH:mm:ssXXX");
  };
  // Email validation function
  const validateEmail = () => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    emailValid = emailRegex.test(email);
  };

  // Contact number validation function
  const validateContactNumber = () => {
    // Add your contact number validation logic here
    // For example, you can use a regular expression or other validation rules
    // For simplicity, let's assume it should be a numeric value with at least 8 digits
    const contactNoRegex = /^\d{8,}$/;
    contactNoValid = contactNoRegex.test(contact_no);
  };

  const fetchUniversities = async () => {
    try {
      const response = await fetch("http://localhost:8080/university/names");
      const data = await response.json();
      universities = data.Message;
      console.log(universities);
    } catch (error) {
      console.error("Error fetching universities:", error);
    }
  };

  const fetchColleges = async () => {
    console.log("fetch colleges", university_id);
    if (university_id) {
      try {
        const response = await fetch(
          `http://localhost:8080/college/names/${university_id}`
        );
        const data = await response.json();
        colleges = data.Message; // Assuming the API returns an array of colleges with properties Id and Title
      } catch (error) {
        console.error("Error fetching colleges:", error);
      }
    }
  };

  // Fetch colleges when the selected university changes
  $: fetchColleges();

  const handleSubmit = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/signup", {
        mode: "no-cors",
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

  onMount(() => {
    // Fetch universities when the component mounts
    fetchUniversities();
  });
</script>

<!-- Remaining HTML and styling unchanged -->
<Header />
<Description />

<main class="text-center max-w-2xl mx-auto p-8">
  <h1 class="text-2xl font-bold mb-4">Create New Account</h1>
  <form on:submit|preventDefault={handleSubmit} class="grid gap-4">
    <label>
      University:
      <select bind:value={university_id} on:change={fetchColleges}>
        <option value="" disabled selected>Select your option</option>
        {#each universities as university}
          <option value={university.ID}>{university.Title}</option>
        {/each}
      </select>
    </label>

    <label>
      College:
      <select bind:value={college_id}>
        <option value="" disabled selected>Select your option</option>
        {#each colleges as college}
          <option value={college.ID}>{college.Title}</option>
        {/each}
      </select>
    </label>

    <label>
      Email:
      <input bind:value={email} type="email" on:input={validateEmail} />
      {#if !emailValid}
        <p class="error">Invalid email address</p>
      {/if}
    </label>

    <label>
      Contact Number:
      <input
        bind:value={contact_no}
        type="tel"
        on:input={validateContactNumber}
      />
      {#if !contactNoValid}
        <p class="error">Invalid contact number</p>
      {/if}
    </label>

    <label>
      Date of Birth:
      <input bind:value={fdob} on:input={formattedDob} type="date" />
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
    /* text-align: center; */
    max-width: 500px;
    margin: 0 auto;
    padding: 2rem;
    font-family: "Khand", sans-serif;
    background-color: #f4f4f4;
  }

  form {
    display: grid;
    gap: 1rem;
  }
</style>

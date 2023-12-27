<script>
  import { onMount } from "svelte";
  import Header from "../lib/Header.svelte";
  import Footer from "../lib/Footer.svelte";
  import Description from "../lib/Description.svelte";
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";

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
  export let data;

  if (browser && data.user_not_signin == true) {
    goto("/signup");
  } else if (browser) {
    goto("/subjects");
  }
</script>

<!-- Remaining HTML and styling unchanged -->
<Header {data} />
<Description />

<Footer />
<!-- 
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
</style> -->

<!-- SubjectPage.svelte -->

<script>
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import SubjectList from "./SubjectList.svelte";
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";
  export let data;
  export let subjects = data.subjects;
  if (browser && data.user_not_signin == true) {
    goto("/signup");
  }
</script>

<!-- Remaining HTML and styling unchanged -->
<Header {data} />
<Description />
<SubjectList {subjects} />
<Footer />

<script>
  import { onMount } from "svelte";
  import Header from "../../../lib/Header.svelte";
  import Footer from "../../../lib/Footer.svelte";
  import Description from "../../../lib/Description.svelte";

  import EmailVerification from "../../../lib/EmailVerification.svelte";
  import { format } from "date-fns";
</script>

<Header />
<Description />
<EmailVerification />
<Footer />

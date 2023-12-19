<!-- SubjectPage.svelte -->

<script>
  import SubjectList from "./SubjectList.svelte";
  import Header from "../../lib/Header.svelte";
  import Footer from "../../lib/Footer.svelte";
  import Description from "../../lib/Description.svelte";
  let subjects = [
    {
      id: "a800b209-d656-4067-bd47-b1cc8a1ce792",
      title: "Mathematics2",
      description: "Study of numbers, quantity, structure, space, and change.",
      university_id: "fdf3a549-5ed7-40f4-8706-9da30265aa7a",
    },
    {
      id: "a800b209-d656-4067-bd47-b1cc8a1ce792",
      title: "Mathematics2",
      description: "Study of numbers, quantity, structure, space, and change.",
      university_id: "fdf3a549-5ed7-40f4-8706-9da30265aa7a",
    },
    {
      id: "a800b209-d656-4067-bd47-b1cc8a1ce792",
      title: "Mathematics2",
      description: "Study of numbers, quantity, structure, space, and change.",
      university_id: "fdf3a549-5ed7-40f4-8706-9da30265aa7a",
    },
    // Add more subjects as needed
  ];
</script>

<!-- Remaining HTML and styling unchanged -->
<Header />
<Description />
<SubjectList {subjects} />
<Footer />

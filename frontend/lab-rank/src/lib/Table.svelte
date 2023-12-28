<script>
  import { onMount } from "svelte";

  // Sample data for initial rendering
  let problems = [];
  export let subjectID;
  export let collegeID;

  onMount(async () => {
    try {
      // Make an API request
      const response = await fetch(
        `http://localhost:8080/problem/${subjectID}/${collegeID}`
      ); // Replace with your API endpoint
      const data = await response.json();

      // Update the problems array with the API response
      console.log(data);
      problems = data.Message || [];

      // Add auto-generated serial numbers
      problems.forEach((problem, index) => {
        problem.serialNumber = index + 1;
      });
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  });

  // Function to handle title click
  const handleTitleClick = async (id, event) => {
    event.preventDefault(); // Prevent the default behavior of the anchor link
    try {
      // Make an API request based on the clicked title's ID
      const response = await fetch(`http://localhost:8080/problem/${id}`);
      const data = await response.json();

      // Log or process the data from the API response
      console.log(data);
    } catch (error) {
      console.error("Error fetching individual problem:", error);
    }
  };
</script>

<main>
  <table>
    <thead>
      <tr>
        <th>Serial Number</th>
        <th style="width: 60%;">Title</th>
        <th style="width: 20%;">Difficulty</th>
      </tr>
    </thead>
    <tbody>
      {#each problems as problem (problem.serialNumber)}
        <tr>
          <td>{problem.serialNumber}</td>
          <td>
            <a href={`/problems/${problem.id}`}>{problem.title}</a>
          </td>
          <td>{problem.difficulty}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</main>

<style>
  table {
    width: 80%;
    border-collapse: collapse;
    margin-top: 20px;
    margin-left: auto;
    margin-right: auto;
    font-family: "Khand", sans-serif;
  }

  th,
  td {
    border: 1px solid #ddd;
    padding: 10px;
    text-align: center;
  }

  th {
    background-color: #f2f2f2; /* Header background color */
  }

  tr:nth-child(even) {
    background-color: #f9f9f9; /* Alternate row background color */
  }

  a {
    text-decoration: none; /* Remove underline from links */
    color: #0366d6; /* Link color */
    cursor: pointer;
  }

  a:hover {
    text-decoration: underline; /* Underline on hover */
  }
</style>

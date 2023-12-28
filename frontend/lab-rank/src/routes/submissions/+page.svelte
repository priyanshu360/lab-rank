<script>
  import Header from "$lib/Header.svelte";
  import Footer from "$lib/Footer.svelte";
  import Description from "$lib/Description.svelte";
  export let data;

  data.submission.forEach((submission, index) => {
    submission.serialNumber = index + 1;
  });
</script>

<Header {data} />
<Description />
<main>
  <table>
    <thead>
      <tr>
        <th>Serial Number</th>
        <th style="width: 60%;">Title</th>
        <th style="width: 20%;">Score</th>
        <th style="width: 20%;">Status</th>
      </tr>
    </thead>
    <tbody>
      {#each data.submission as submission (submission.serialNumber)}
        <tr>
          <td>{submission.serialNumber}</td>
          <td>
            <a href={`/submissions/${submission.id}`}>{submission.problem_id}</a
            >
          </td>
          <td>{submission.score}</td>
          <td>{submission.status}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</main>
<Footer />

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

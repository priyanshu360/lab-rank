<script>
  let title = "LabRank";
  import { goto } from "$app/navigation";
  const handleSignInClick = () => {
    // Navigate to the "/login" page
    console.log("click");
    goto("/login");
  };
  export let data;
</script>

<header class="bg-white">
  <nav class="flex justify-between items-center w-[92%] mx-auto">
    <div>
      <a href="/">{title}</a>
    </div>
    <div
      class="nav-links duration-500 md:static absolute md:min-h-fit min-h-[60vh] left-0 top-[-100%] md:w-auto w-full flex items-center px-5"
    >
      <ul class="flex md:flex-row flex-col md:items-center md:gap-[4vw] gap-8">
        <li>
          <a class="hover:text-gray-500" href="/problems">Problems</a>
        </li>
        <li>
          <a class="hover:text-gray-500" href="/subjects">Subjects</a>
        </li>
        <li>
          <a class="hover:text-gray-500" href="/submissions">Submissions</a>
        </li>
      </ul>
    </div>
    <div class="flex items-center gap-6">
      {#if data.user_not_signin}
        <button
          on:click={handleSignInClick}
          href
          class="bg-[#a6c1ee] text-white px-5 py-2 rounded-full hover:bg-[#87acec]"
          >Sign in</button
        >
      {:else}
        <button
          on:click={handleSignInClick}
          href
          class="bg-[#a6c1ee] text-white px-5 py-2 rounded-full hover:bg-[#87acec]"
          >Sign Out</button
        >
      {/if}
      <ion-icon
        onclick="onToggleMenu(this)"
        name="menu"
        class="text-3xl cursor-pointer md:hidden"
      ></ion-icon>
    </div>
  </nav>
</header>

<style>
  header {
    background-color: #ffffff; /* Light background color */
    color: #0c0c0c; /* Dark text color */
    text-align: left;
    width: 100%;
    font-family: "Khand", sans-serif;
  }

  a {
    text-decoration: none;
  }
</style>

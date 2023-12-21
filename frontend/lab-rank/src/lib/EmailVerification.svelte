<script>
  let otp = ["", "", "", ""];
  let verificationStatus = "Enter OTP";

  // Function to handle OTP input
  const handleOtpInput = (index, event) => {
    const value = event.target.value;

    // Ensure the input is a single digit
    if (/^\d$/.test(value)) {
      otp[index] = value;

      // Move to the next input field if available
      if (index < otp.length - 1 && value !== "") {
        inputFields[index + 1]?.focus();
      }
    }
  };

  // Array to store references to input fields
  let inputFields = [];

  // Simulate a backend call for OTP verification
  const verifyOtp = async () => {
    try {
      const otpValue = otp.join("");

      // Simulate an API call to your backend for OTP verification
      // Replace this with an actual API call to your server
      const response = await fetch(
        `/api/verify-otp?email=${email}&otp=${otpValue}`
      );
      const data = await response.json();

      // Assuming the server returns a success message upon successful verification
      if (data.success) {
        verificationStatus = "OTP Verified Successfully!";
      } else {
        verificationStatus = "OTP Verification Failed";
      }
    } catch (error) {
      console.error("Error during OTP verification:", error);
      verificationStatus = "Error during OTP verification";
    }
  };
</script>

<main class="text-center max-w-2xl mx-auto p-8">
  <h1>OTP Verification</h1>
  <form on:submit|preventDefault={verifyOtp}>
    {#each otp as digit, index (index)}
      <input
        type="text"
        value={digit}
        on:input={(e) => handleOtpInput(index, e)}
        bind:this={inputFields[index]}
        class="otp-input"
        maxlength="1"
      />
    {/each}
    <button
      type="submit"
      class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded"
      >Verify OTP</button
    >
  </form>
  <p>{verificationStatus}</p>
</main>

<style>
  .otp-input {
    width: 2em;
    margin: 0.05em;
    text-align: center;
    border-color: #0c0c0c;
    border-width: 0.01em;
  }
  main {
    background-color: #f4f4f4; /* Light background color */
    color: #0c0c0c; /* Dark text color */
    /* text-align: left; */
    width: 100%;
    font-family: "Khand", sans-serif;
  }
</style>

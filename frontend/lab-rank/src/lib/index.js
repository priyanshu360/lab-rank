// place files you want to import through the `$lib` alias in this folder.

export const makeLoginReq = async (email, password) => {
  try {
    // Check if the server is reachable
    const isServerReachable = await isReachable("http://localhost:8080");
    
    if (!isServerReachable) {
      throw new Error("Server is not reachable");
    }

    const res = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email,
        password,
      }),
    });

    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    return res.json().Message;
  } catch (error) {
    console.error("Error during login:", error);
    throw error; // Propagate the error
  }
};

const isReachable = async (url) => {
  try {
    await fetch(url, { method: 'HEAD' });
    return true;
  } catch {
    return false;
  }
};
// sessionStore.js
import { writable } from 'svelte/store';

export const sessionStore = writable({
  subjects: [], // Initial subjects array
  // Add other properties as needed
});

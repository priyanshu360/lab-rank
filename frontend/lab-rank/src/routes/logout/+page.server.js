import { redirect } from '@sveltejs/kit';

export const load = ({ cookies }) => {
    cookies.delete("jwt-lab-rank")
    cookies.delete("college-id")
    cookies.delete("university-id")
    throw redirect(303, '/login')
}
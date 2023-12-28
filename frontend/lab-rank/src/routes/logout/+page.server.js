import { redirect } from '@sveltejs/kit';

export const load = ({ cookies }) => {
    cookies.delete("jwt-lab-rank")
    throw redirect(303, '/login')
}
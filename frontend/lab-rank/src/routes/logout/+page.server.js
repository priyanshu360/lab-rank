import { redirect } from '@sveltejs/kit';

export const load = ({ cookies }) => {
    cookies.delete("jwt_lab_rank")
    throw redirect(303, '/login')
}
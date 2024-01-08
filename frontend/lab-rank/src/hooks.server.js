/** @type {import('@sveltejs/kit').Handle} */
/** @type {import('@sveltejs/kit').HandleFetch} */
import cookie from 'cookie';

import { redirect } from '@sveltejs/kit';


async function authenticate(jwtToken) {
    // Make a request to your backend API for authentication
    const response = await fetch('http://127.0.0.1:8080/auth/user', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${jwtToken}`
        },
        // Add any necessary body parameters if required
        // body: JSON.stringify({ /* your data */ }),
    });

    // Check if the authentication was successful
    if (!response.ok) {
        // You can handle the error here, for example, throw an exception
        throw new Error('Authentication failed');
    }

    // Return the response or parsed data if needed
    return response.json();
}


export async function handle({ event, resolve }) {
    const { request } = event;
    const cookies = cookie.parse(request.headers.get('cookie') || '');
    const jwtToken = cookies.jwt_lab_rank;

    if (!jwtToken && !(event.url.pathname.startsWith('/sign') || event.url.pathname.startsWith('/log'))) {
        throw redirect(303, '/signup')
    } else if (jwtToken) {
        try {
            console.log("found jwt")
            const res = await authenticate(jwtToken);
            // console.log("res", res.Message)
            event.locals.user = res.Message
        } catch (error) {
            // Handle authentication error (e.g., redirect to /login)
            console.log(error)
            // Todo : delete coookie
            // cookie.delete()
            throw redirect(303, '/signup');
        }
    }

    // console.log("jwt verified", event)


    const response = await resolve(event); return response;
}


export async function handleFetch({ event, request, fetch }) {
    console.log("hello my dear friend")
    const cookies = cookie.parse(event.request.headers.get('cookie') || '');
    const jwtToken = cookies.jwt_lab_rank;
    console.log(cookies, jwtToken)
    if (jwtToken) {
        // clone the original request, but change the URL
        request.headers.set('Authorization', `Bearer ${jwtToken}`)
    }

    console.log(request)

    return fetch(request);
}
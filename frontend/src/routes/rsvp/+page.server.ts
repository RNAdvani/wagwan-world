import type { Actions, PageServerLoad } from './$types';
import {env} from '$env/dynamic/private';
import {fail} from '@sveltejs/kit';

const API_BASE_URL = 'http://localhost:8080/api';

export const load: PageServerLoad = async ({fetch}) => {
    const response = await fetch(`${API_BASE_URL}/events`);
    if (!response.ok) {
        throw new Error('Failed to load events');
    }
    const events = await response.json();
    return { events };
}


export const actions: Actions = {
    default: async ({request, fetch}) => {
        const formData = await request.formData();

        const payload = {
            name: String(formData.get('name') ?? ''),
            email: String(formData.get('email') ?? ''),
            phone: String(formData.get('phone') ?? ''),
            status: String(formData.get('status') ?? 'pending'),
            event_id: Number(formData.get('event_id') ?? 1),
            notes: String(formData.get('notes') ?? ''),
            plus_ones: Number(formData.get('plus_ones') ?? 0),
            dietary_restrictions: String(formData.get('dietary_restrictions') ?? '')
        };

        if (!payload.name || !payload.email) {
            return fail(400, { formError: 'Name and email are required' });
        }

        const res = await fetch(`${API_BASE_URL}/guests`, {
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify(payload)
        });

        if (!res.ok) {
            const err = await res.json().catch(() => ({ error: 'Server error' }));
            console.log(res.status)
            return fail(res.status, { formError: err.error ?? 'Failed to submit RSVP' });
        }

    
        return { success: true };
    }
}
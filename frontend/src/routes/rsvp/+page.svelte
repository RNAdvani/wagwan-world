<script lang="ts">
	import { onMount } from 'svelte';
	import { getEvents, createGuest, type Event as EventType } from '$lib/api';

	let events: EventType[] = [];
	let loading = true;
	let error: string | null = null;

	let formName = '';
	let formEmail = '';
	let formPhone = '';
	let formStatus: 'attending' | 'pending' | 'declined' = 'pending';
	let formEventId: number | null = null;
	let formNotes = '';
	let formPlusOnes: 0 | 1 = 0;
	let formDietaryRestrictions = '';

	let formError: string | null = null;
	let formSubmitting = false;
	let formSuccess = false;

	onMount(async () => {
		try {
			events = await getEvents();
		} catch (err) {
			error = 'Failed to load events';
		} finally {
			loading = false;
		}
	});

	async function handleSubmit(e: Event) {
		e.preventDefault();
		formError = null;
		formSuccess = false;

		if (!formEventId) {
			formError = 'Please select an event';
			return;
		}
		if (!formName.trim()) {
			formError = 'Name is required';
			return;
		}
		if (!formEmail.trim()) {
			formError = 'Email is required';
			return;
		}
		if (!/\S+@\S+\.\S+/.test(formEmail)) {
			formError = 'Invalid email format';
			return;
		}
		if (!formPhone.trim()) {
			formError = 'Phone is required';
			return;
		}

		try {
			formSubmitting = true;

			await createGuest({
				name: formName,
				email: formEmail,
				phone: formPhone,
				status: formStatus,
				event_id: formEventId,
				notes: formNotes,
				plus_ones: formPlusOnes,
				dietary_restrictions: formDietaryRestrictions
			});

			formName = '';
			formEmail = '';
			formPhone = '';
			formStatus = 'pending';
			formEventId = null;
			formNotes = '';
			formPlusOnes = 0;
			formDietaryRestrictions = '';
			formSuccess = true;
		} catch (err) {
			formError =
				err instanceof Error ? err.message : 'Failed to submit RSVP';
		} finally {
			formSubmitting = false;
		}
	}

	function formatDate(date: string) {
		return new Date(date).toLocaleString('en-IN', {
			day: 'numeric',
			month: 'long',
			year: 'numeric',
			hour: '2-digit',
			minute: '2-digit'
		});
	}
</script>

<div class="min-h-screen  bg-indigo-50">
	<div class="container mx-auto px-4 py-12 max-w-6xl">
		<div class="mb-8 bg-white/80 rounded-xl border border-gray-200 p-4">
			<h1 class="text-2xl font-bold text-gray-900 mb-2">RSVP to Our Events</h1>
			<p class="text-gray-600">Select an event below and let us know you're coming</p>
		</div>


		{#if loading}
			<div class="flex justify-center items-center h-64">
				<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
			</div>
		{:else if error}
			<p class="text-center text-red-500 text-lg">{error}</p>
		{:else}
			<div class="grid lg:grid-cols-3 gap-8">
				<!-- Events Sidebar - 1/3 width -->
				<div class="lg:col-span-1">
					<div class="sticky top-8">
						<div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl border border-gray-100 p-6">
							<h2 class="text-2xl font-bold text-gray-900 mb-4">
								Events
							</h2>
							<div class="space-y-3 overflow-y-auto pr-2 styled-scroll" style="max-height: 500px;">
								{#each events as e}
									<div
										on:click={() => (formEventId = e.id)}
										on:keydown={(ev) => ev.key === 'Enter' && (formEventId = e.id)}
										role="button"
										tabindex="0"
										class="group cursor-pointer border-2 rounded-2xl p-4 transition-all duration-200
											{formEventId === e.id 
												? 'border-blue-600 bg-blue-50 shadow-lg' 
												: 'border-gray-200 bg-white hover:border-blue-400 hover:shadow-md'}">
										<h3 class="font-bold text-gray-900 mb-1 text-sm group-hover:text-blue-600 transition-colors">
											{e.title}
										</h3>
										<p class="text-gray-500 text-xs mb-2 line-clamp-2">{e.description}</p>
										<div class="text-xs text-gray-600 space-y-1">
											<p class="truncate">{formatDate(e.event_date)}</p>
											<p class="truncate">{e.location}</p>
										</div>
									</div>
								{/each}
							</div>
						</div>
					</div>
				</div>

				<!-- RSVP Form - 2/3 width -->
				<div class="lg:col-span-2">
					<div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl border border-gray-100 p-8">
						<h2 class="text-3xl font-bold text-gray-900 mb-6">
							Your Details
						</h2>
						
						<form class="space-y-5" on:submit|preventDefault={handleSubmit}>
							<div class="grid md:grid-cols-2 gap-5">
								<!-- Name -->
								<div>
									<p class="text-sm font-bold text-gray-700 mb-2">Full Name</p>
									<input
										type="text"
										bind:value={formName}
										class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:border-blue-600 bg-white transition-all"
										placeholder="John Doe" />
								</div>

								<!-- Email -->
								<div>
									<p class="text-sm font-bold text-gray-700 mb-2">Email Address</p>
									<input
										type="email"
										bind:value={formEmail}
										class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:border-blue-600 bg-white transition-all"
										placeholder="john@example.com" />
								</div>

								<!-- Phone -->
								<div>
									<p class="text-sm font-bold text-gray-700 mb-2">Phone Number</p>
									<input
										type="tel"
										bind:value={formPhone}
										class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:border-blue-600 bg-white transition-all"
										placeholder="+91 98765 43210" />
								</div>

								<!-- Plus Ones -->
								<div>
									<p class="text-sm font-bold text-gray-700 mb-2">Bringing a Guest?</p>
									<select
										bind:value={formPlusOnes}
										class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:border-blue-600 bg-white transition-all cursor-pointer">
										<option value={0}>Just me</option>
										<option value={1}>Me + 1</option>
									</select>
								</div>
							</div>

							<!-- Dietary Restrictions -->
							<div>
								<p class="text-sm font-bold text-gray-700 mb-2">Dietary Preferences</p>
								<input
									type="text"
									bind:value={formDietaryRestrictions}
									class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:border-blue-600 bg-white transition-all"
									placeholder="Vegetarian, Vegan, None, etc." />
							</div>

							<!-- Notes -->
							<div>
								<p class="text-sm font-bold text-gray-700 mb-2">Notes</p>
								<textarea
									rows="3"
									bind:value={formNotes}
									class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:border-blue-600 bg-white transition-all resize-none"
									placeholder="Any special requests or messages for us?"></textarea>
							</div>

							<!-- RSVP Status -->
							<div>
								<p class="text-sm font-bold text-gray-700 mb-3">Will you be attending?</p>
								<div class="flex flex-wrap gap-3">
									<div 
										on:click={() => formStatus = 'attending'}
										on:keydown={(ev) => ev.key === 'Enter' && (formStatus = 'attending')}
										role="button"
										tabindex="0"
										class="flex items-center gap-2 px-5 py-3 rounded-lg border-2 cursor-pointer transition-all
											{formStatus === 'attending' ? 'border-green-500 bg-green-50' : 'border-gray-300 bg-white hover:border-green-400'}">
										<input class="cursor-pointer w-4 h-4 pointer-events-none" type="radio" bind:group={formStatus} value="attending" />
										<span class="font-medium">Yes, I'll be there!</span>
									</div>
									<div 
										on:click={() => formStatus = 'pending'}
										on:keydown={(ev) => ev.key === 'Enter' && (formStatus = 'pending')}
										role="button"
										tabindex="0"
										class="flex items-center gap-2 px-5 py-3 rounded-xl border-2 cursor-pointer transition-all
											{formStatus === 'pending' ? 'border-yellow-500 bg-yellow-50' : 'border-gray-300 bg-white hover:border-yellow-400'}">
										<input class="cursor-pointer w-4 h-4 pointer-events-none" type="radio" bind:group={formStatus} value="pending" />
										<span class="font-medium">Maybe</span>
									</div>
									<div 
										on:click={() => formStatus = 'declined'}
										on:keydown={(ev) => ev.key === 'Enter' && (formStatus = 'declined')}
										role="button"
										tabindex="0"
										class="flex items-center gap-2 px-5 py-3 rounded-xl border-2 cursor-pointer transition-all
											{formStatus === 'declined' ? 'border-red-500 bg-red-50' : 'border-gray-300 bg-white hover:border-red-400'}">
										<input class="cursor-pointer w-4 h-4 pointer-events-none" type="radio" bind:group={formStatus} value="declined" />
										<span class="font-medium">Sorry, can't make it</span>
									</div>
								</div>
							</div>

							{#if formError}
								<div class="bg-red-50 border-2 border-red-200 rounded-xl p-4">
									<p class="text-sm text-red-600 text-center font-medium">{formError}</p>
								</div>
							{/if}
							{#if formSuccess}
								<div class="bg-green-50 border-2 border-green-200 rounded-xl p-4">
									<p class="text-sm text-green-600 text-center font-medium">RSVP submitted successfully!</p>
								</div>
							{/if}

							<!-- Submit Button -->
							<button
								type="submit"
								class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-4 rounded-xl transition-all disabled:opacity-50 shadow-lg hover:shadow-xl transform hover:-translate-y-0.5"
								disabled={formSubmitting}>
								{#if formSubmitting}
									<span class="flex items-center justify-center gap-2">
										<span class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></span>
										<span>Submitting...</span>
									</span>
								{:else}
									<span class="text-lg">Submit RSVP</span>
								{/if}
							</button>
						</form>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>

<style>
	.styled-scroll::-webkit-scrollbar {
		width: 6px;
	}

	.styled-scroll::-webkit-scrollbar-track {
		background: transparent;
	}

	.styled-scroll::-webkit-scrollbar-thumb {
		background: linear-gradient(to bottom, #818cf8, #a78bfa);
		border-radius: 10px;
	}

	.styled-scroll::-webkit-scrollbar-thumb:hover {
		background: linear-gradient(to bottom, #6366f1, #9333ea);
	}

	.line-clamp-2 {
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}
</style>
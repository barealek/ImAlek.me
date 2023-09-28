<script>
	import About from '$lib/components/About.svelte';
	import Resume from '$lib/components/Resume.svelte';
	import Portfolio from '$lib/components/Portfolio.svelte';
	import '../page.css';


	const pages = ['about', 'resume', 'portfolio']
	let page = pages[0];

	let expandSidebar = false;
</script>

<svelte:head>
	<title>ImAlek - Software Developer</title>
</svelte:head>

<!--
    - #MAIN
  -->
<main>
	<!--
      - #SIDEBAR
    -->

	<aside class="sidebar {expandSidebar ? 'active' : ''}">
		<div class="sidebar-info">
			<figure class="avatar-box">
				<img class="rounded-2xl" src="/assets/me.webp" alt="Profile pic" width="80" height="80" />
			</figure>

			<div class="info-content">
				<h1 class="text-center text-white text-xl mb-3 font-semibold" title="ImAlek">ImAlek</h1>

				<p class="text-white rounded-lg px-3 py-1 text-xs w-max text-center font-medium bg-[#4073b7]">Software Developer</p>
			</div>

			<button class="info_more-btn" on:click={() => {expandSidebar = !expandSidebar}}>
				<span>Show Contacts</span>

				<ion-icon name="chevron-down" />
			</button>
		</div>

		<div class="sidebar-info_more active">
			<div class="separator" />

			<ul class="contacts-list">
				<li class="contact-item">
					<div class="icon-box">
						<ion-icon name="mail-outline" />
					</div>

					<div class="contact-info">
						<p class="contact-title">Email</p>

						<a href="mailto:alek@imalek.me" class="contact-link">alek@imalek.me</a>
					</div>
				</li>

				<li class="contact-item">
					<div class="icon-box">
						<ion-icon name="logo-discord" />
					</div>

					<div class="contact-info">
						<p class="contact-title">Discord</p>
						<a href="https://discord.com/users/424265745627873302" class="contact-link"
							>alek.dev</a
						>
					</div>
				</li>

				<li class="contact-item">
					<div class="icon-box">
						<ion-icon name="logo-github" />
					</div>

					<div class="contact-info">
						<p class="contact-title">GitHub</p>

						<a href="https://github.com/barealek" target="_blank" class="contact-link">@BareAlek</a>
					</div>
				</li>

				<li class="contact-item">
					<div class="icon-box">
						<ion-icon name="location-outline" />
					</div>

					<div class="contact-info">
						<p class="contact-title">Location</p>

						<address>Denmark</address>
					</div>
				</li>
			</ul>

			<div class="separator" />
		</div>
	</aside>

	<!--
      - #main-content
    -->

	<div class="main-content">

		<nav class="navbar">
			<ul class="navbar-list">
				{#each pages as p}
				<li class="navbar-item">
					<button on:click={() => {page = p}} class="navbar-link {page == p ? 'active' : ''} font-normal capitalize">{p}</button>
				</li>
				{/each}
			</ul>
		</nav>

		{#if page === 'about'}
		<About/>

		{:else if page === 'resume'}

		<Resume/>

		{:else if page === 'portfolio'}
		
		<Portfolio/>
		
		{/if}
	</div>
</main>

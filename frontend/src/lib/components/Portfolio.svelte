<script>
    let projects = [
        {
            title: "miniLink - URL Shortener",
            category: "REST API | Docker",
            img: "/assets/minilink.png",
            link: "https://github.com/barealek/miniLink"
        },
        {
            title: "StyledLogger - Logging Lib",
            category: "Python | CI/CD",
            img: "/assets/styledlogger.png",
            link: "https://github.com/barealek/styledlogger"
        },
        {
            title: "ImAlek.me - Portfolio",
            category: "Auth | REST API | Docker",
            img: "/assets/homepage.png",
            link: "https://github.com/barealek/imalek.me"
        }
    ]
</script>
<svelte:head>
    <title>ImAlek - Portfolio</title>
</svelte:head>

<article>
    <header>
        <h2 class="text-white text-3xl mb-1 font-bold pb-4 after:my-4 after:block after:w-10 after:h-1 after:bg-[#4797ff] after:rounded-lg">Portfolio</h2>
    </header>

    <section>

        <ul class="grid gap-6 mb-4 grid-cols-3">
            {#each projects as project, i (i)}
            <li class="animate-fade-down animate-once animate-duration-500 animate-delay-[{150*i}ms] animate-ease-in-out">
                <a href={project.link} target="_blank">
                    <figure class="relative w-[100%] h-[150px] rounded-xl overflow-hidden md:mb-1">
                        <img src={project.img} alt={project.title} loading="lazy" class="md:aspect-video hover:scale-105 transition-transform" />
                    </figure>

                    <h3 class="text-white text-md ml-2">{project.title}</h3>
                    <p class="text-gray-500 text-sm ml-2">{project.category}</p>
                </a>
            </li>
            {/each}
        </ul>
    </section>
</article>
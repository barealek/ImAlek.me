<svelte:head>
    <title>ImAlek - Resume</title>
</svelte:head>

<script>
    let experiences = [
        {
            title: "Jet-Market.dk [MC Marketplace]",
            role: "Founder & Developer",
            date: "Sep 2023 — Present",
            description: [
                "• API Developer",
                "• Frontend Developer",
                "• Service Management",
                "• Infrastructure Management"
            ]
        },
        {
            title: "Jet-Host.dk [Hosting Firm]",
            role: "Co-Founder",
            date: "Aug 2023 — Present",
            description: [
                "• Workflow Developer",
                "• Infrastructure Management",
                "• Community Engagement",
                "• Moral Support <3"
            ]
        },
        {
            title: "AntiToxicity [Discord Bot]",
            role: "Sole Developer",
            date: "Oct 2022 — Nov 2022",
            description: [
                "• Project Structure (at 14)",
                "• Infrastructure Management",
                "• Website Development (w/ help, as I was 14)",
                "• Machine Learning Toxicity Pipeline",
                "⚠️This project is rather old and doesn't reflect my current experience."
            ]
        },
        {
            title: "Open Source",
            role: "Contributions",
            date: "",
            description: [
                "• barealek/StyledLogger",
                "• nwunderly/Starlette-Discord",
                "• Lukkit/Lukkit",
                "• And more!"
            ]
        },
        {
            title: "Freelance Work",
            role: "Projects",
            date: "",
            description: [
                "• Discord Bot @ .gg/Worldwide",
                "• Discord Bot @ .gg/TheLake",
                "• Website for a few local companies",
                "• And much more!"
            ]
        }
    ]
</script>


<article>
    <header>
        <h2 class="text-white text-3xl mb-1 font-bold pb-4 after:my-4 after:block after:w-10 after:h-1 after:bg-[#4797ff] after:rounded-lg">Resume</h2>
    </header>

    <section class="mb-8">
        <div class="flex items-center gap-4 mb-6 ">
            <div class="icon-box">
                <ion-icon name="book-outline" />
            </div>

            <h3 class="h3">Experience</h3>
        </div>

        <ol class="ml-16 text-blue-500 text-sm">
            {#each experiences as exp, i (i) }
            <li class="timeline-item animate-fade-up animate-duration-300 animate-delay-[{150*i}ms] animate-ease-in-out">
                <h4 class="text-md text-white">{exp.title}</h4>
                <span>{exp.date !== "" ? [exp.role, exp.date].join(" | ") : exp.role}</span>

                {#each exp.description as desc}
                    <p class="text-gray-400">{desc}</p>
                {/each}
            </li>
            {/each}
        </ol>
    </section>
</article>
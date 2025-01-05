<script lang="ts">
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeProjectsSectionItemCarousel
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemCarousel/HomeProjectsSectionItemCarousel.svelte";
    import type {BestService} from "@/types/BestService.ts";

    type HomeProjectsSectionProps = {
        bestServices: BestService[],
    };

    const props: HomeProjectsSectionProps = $props();
</script>

<HomeSectionLayout fullHeight>
    <HomeSectionHeader
            title="Featured Projects"
            description="Explore some of the incredible projects completed by our talented freelancers."
    />
    <HomeProjectsSectionItemCarousel {...props}/>
</HomeSectionLayout>

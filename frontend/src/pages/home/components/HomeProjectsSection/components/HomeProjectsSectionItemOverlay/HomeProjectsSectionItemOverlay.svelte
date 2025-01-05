<script lang="ts">
    import HomeProjectsSectionItemOverlayLayout
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemOverlay/HomeProjectsSectionItemOverlayLayout.svelte";
    import HomeProjectsSectionItemOverlayDescription
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemOverlayDescription/HomeProjectsSectionItemOverlayDescription.svelte";
    import HomeProjectsSectionItemOverlayIcon
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemOverlayIcon/HomeProjectsSectionItemOverlayIcon.svelte";
    import HomeProjectsSectionItemOverlayTitle
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemOverlayTitle/HomeProjectsSectionItemOverlayTitle.svelte";
    import HomeProjectsSectionItemOverlayContent
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemOverlayContent/HomeProjectsSectionItemOverlayContent.svelte";
    import type {BestService} from "@/types/BestService.ts";
    import Button from "@/common-components/Button/Button.svelte";
    import HomeProjectsItemOverlayButton
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsItemOverlayButton/HomeProjectsItemOverlayButton.svelte";

    const props: BestService = $props();
</script>

<HomeProjectsSectionItemOverlayLayout>
    <HomeProjectsSectionItemOverlayContent>
        <HomeProjectsSectionItemOverlayIcon {...props}/>
        <HomeProjectsSectionItemOverlayTitle {...props}/>
    </HomeProjectsSectionItemOverlayContent>
    <HomeProjectsSectionItemOverlayDescription {...props}/>
    <HomeProjectsItemOverlayButton {...props}/>
</HomeProjectsSectionItemOverlayLayout>
<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import HomeProjectsSectionItemImage
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemImage/HomeProjectsSectionItemImage.svelte";
    import HomeProjectsSectionItemOverlay
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItemOverlay/HomeProjectsSectionItemOverlay.svelte";
    import type {BestService} from "@/types/BestService.ts";

    const props: BestService = $props();
</script>

<PaperElement styles="relative w-full !p-0 !ring-0 !shadow-none !bg-transparent">
    <HomeProjectsSectionItemImage {...props}/>
    <HomeProjectsSectionItemOverlay {...props}/>
</PaperElement>
<script lang="ts">
    import Button from "@/common-components/Button/Button.svelte";
    import HomeProjectsItemOverlayButtonLayout
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsItemOverlayButton/HomeProjectsItemOverlayButtonLayout.svelte";
    import {Link} from "svelte-routing";
    import type {BestService} from "@/types/BestService.ts";

    const {id} : BestService = $props();
</script>

<HomeProjectsItemOverlayButtonLayout>
    <Link to="/services/{id}">
        <Button styles="w-fit">Learn More</Button>
    </Link>
</HomeProjectsItemOverlayButtonLayout>

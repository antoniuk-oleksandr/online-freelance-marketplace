<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import HomeReviewsSectionItemContent
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItemContent/HomeReviewsSectionItemContent.svelte";
    import HomeReviewsSectionItemName
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItemName/HomeReviewsSectionItemName.svelte";
    import HomeReviewsSectionItemRating
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItemRating/HomeReviewsSectionItemRating.svelte";
    import type {BestReview} from "@/types/BestReview.ts";

    const props: BestReview = $props();
</script>

<PaperElement styles="flex flex-col gap-y-2">
    <HomeReviewsSectionItemRating {...props}/>
    <HomeReviewsSectionItemContent {...props}/>
    <HomeReviewsSectionItemName {...props}/>
</PaperElement>
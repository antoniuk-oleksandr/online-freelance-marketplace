<script lang="ts">
    import Icon from "@iconify/svelte";
    import HomeReviewsSectionItemRatingLayout
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItemRating/HomeReviewsSectionItemRatingLayout.svelte";

    type HomeReviewsSectionItemRatingProps = {
        rating: number,
    };

    const {rating}: HomeReviewsSectionItemRatingProps = $props();
</script>

<HomeReviewsSectionItemRatingLayout>
    {#each Array(rating) as _}
        <Icon
                icon="mingcute:star-fill"
                class="text-yellow-500 dark:text-yellow-500"
                width="24" height="24"
        />
    {/each}
</HomeReviewsSectionItemRatingLayout>
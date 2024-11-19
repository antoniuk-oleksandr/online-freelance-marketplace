<script lang="ts">
    import UserByIdReviewRatingBarLayout
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewRatingBar/UserByIdReviewRatingBarLayout.svelte";
    import Rating from "@/common-components/Rating/Rating.svelte";
    import type {Review} from "@/types/Review.ts";
    import BulletElement from "@/common-components/BulletElement/BulletElement.svelte";
    import {getTimeAgo} from "@/utils/utils.ts";

    const {rating, endedAt}: Review = $props();
</script>

<UserByIdReviewRatingBarLayout>
    <Rating value={rating}/>
    <p class="text-base font-semibold">{rating}</p>
    <BulletElement/>
    <p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
    >{getTimeAgo(endedAt)}</p>
</UserByIdReviewRatingBarLayout>



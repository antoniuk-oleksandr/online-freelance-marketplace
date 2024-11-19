<script lang="ts">
    import type {UserReview} from "@/types/UserReview.ts";
    import UserByIdReviewServiceLayout
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewService/UserByIdReviewServiceLayout.svelte";
    import UserByIdReviewServiceImage
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewServiceImage/UserByIdReviewServiceImage.svelte";

    const {image, title, id}: UserReview = $props();
</script>

<UserByIdReviewServiceLayout id={id}>
    <UserByIdReviewServiceImage src={image}/>
    <span class="line-clamp-2 w-auto md:w-40">{title}</span>
</UserByIdReviewServiceLayout>
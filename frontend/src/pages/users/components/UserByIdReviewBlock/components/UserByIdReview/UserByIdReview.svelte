<script lang="ts">
    import type {Review} from "@/types/Review.ts";
    import UserByIdReviewLayout
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/UserByIdReviewLayout.svelte";
    import UserByIdReviewUser
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewUser/UserByIdReviewUser.svelte";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import UserByIdReviewRatingBar
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewRatingBar/UserByIdReviewRatingBar.svelte";
    import UserByIdReviewBottomContent
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewBottomContent/UserByIdReviewBottomContent.svelte";
    import UserByIdReviewContent
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewContent/UserByIdReviewContent.svelte";

    type UserByIdReviewProps = {
        review: Review,
        showServices?: boolean
    }

    const {review, showServices}:  UserByIdReviewProps = $props();
</script>

<UserByIdReviewLayout>
    <UserByIdReviewUser {...review.customer}/>
    <DividerElement/>
    <UserByIdReviewRatingBar {...review}/>
    <UserByIdReviewContent {...review}/>
    <UserByIdReviewBottomContent
            review={review}
            showServices={showServices}
    />
</UserByIdReviewLayout>
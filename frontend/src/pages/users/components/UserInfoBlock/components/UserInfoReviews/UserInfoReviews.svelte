<script lang="ts">
    import Icon from "@iconify/svelte";
    import UserInfoReviewsLayout
        from "@/pages/users/components/UserInfoBlock/components/UserInfoReviews/UserInfoReviewsLayout.svelte";
    import type {UserByIdProps} from "@/types/UserByIdProps.ts";

    type UserInfoReviewsProps = UserByIdProps & {
        size?: "small" | "large",
    };

    const {user, size}: UserInfoReviewsProps = $props();
</script>

{#if user.rating !== null && user.reviewsCount !== null}
    <UserInfoReviewsLayout size={size}>
        <Icon class="text-yellow-500" icon="mingcute:star-fill"/>
        <span class="font-semibold">{user.rating}</span>
        <span>({user.reviewsCount})</span>
        <span>Level {user.level}</span>
    </UserInfoReviewsLayout>
{/if}
<script lang="ts">
    import UserByIdReviewColumnLayout
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewColumn/UserByIdReviewColumnLayout.svelte";

    type UserByIdReviewColumnProps = {
        label: string,
        data: string | number
    }

    const {label, data}: UserByIdReviewColumnProps = $props();
</script>

<UserByIdReviewColumnLayout>
    <p class="font-semibold text-base">{data}</p>
    <p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary w-20 md:w-fit">{label}</p>
</UserByIdReviewColumnLayout>
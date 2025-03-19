<script lang="ts">
    import UserByIdReviewServiceImageLayout
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewServiceImage/UserByIdReviewServiceImageLayout.svelte";

    type UserByIdReviewServiceImageProps = {
        src: string
    }

    let {src}: UserByIdReviewServiceImageProps = $props();
    if(!src) src = "/images/no-image.ong";
</script>

<UserByIdReviewServiceImageLayout>
    <img
            width={48}
            height={48}
            src={src}
            alt="img"
            class="object-center aspect-video w-full object-cover h-9"
    />
</UserByIdReviewServiceImageLayout>

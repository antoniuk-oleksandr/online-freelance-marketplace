<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";

    let {image, title}: UserService = $props();

    let imageSrc = $state<string | undefined>();

    $effect(() => {
        imageSrc = image ? image : '/images/no-image.png';
    })
</script>

{#if imageSrc}
    <img
            src={imageSrc}
            alt={title}
            class="object-cover !aspect-video !rounded-md mb-3 size-full"
    />
{/if}

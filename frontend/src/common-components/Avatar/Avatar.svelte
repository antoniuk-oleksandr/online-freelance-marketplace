<script lang="ts">
    import Icon from "@iconify/svelte";

    type AvatarProps = {
        src: string | null,
        alt: string,
        size?: "extra small" | "small" | "large",
    }

    const {src, alt, size}: AvatarProps = $props();

    let sizeClass = $state("");
    switch (size){
        case "extra small":
            sizeClass = "size-12";
            break;
        case "small":
            sizeClass = "size-16";
            break;
        case "large":
            sizeClass = "size-32";
            break;
    }

</script>

<div class="bg-light-palette-background-block dark:bg-dark-palette-background-block rounded-full">
    {#if src}
        <img
                class={`rounded-full object-center object-cover ${sizeClass}`}
                src={src} alt={alt}
        >
    {:else}
        <Icon
                class={`rounded-full text-cyan-500 ${sizeClass}`}
                icon="hugeicons:user-circle-02"
        />
    {/if}
</div>
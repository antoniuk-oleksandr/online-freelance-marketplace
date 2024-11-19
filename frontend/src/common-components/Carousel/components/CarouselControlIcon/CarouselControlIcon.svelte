<script lang="ts">
    import Icon from "@iconify/svelte";

    type CarouselControlIconProps = {
        clickAction: () => void,
        direction: "left" | "right"
    };

    const {clickAction, direction}: CarouselControlIconProps = $props();
</script>

<button
        class="bg-black text-white rounded-full p-2 shadow-md hover:bg-gray-700 duration-200 ease-out"
        onclick={clickAction}
>
    <Icon icon="material-symbols:chevron-{direction}"/>
</button>
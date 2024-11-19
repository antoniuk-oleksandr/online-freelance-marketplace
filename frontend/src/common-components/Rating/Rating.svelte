<script lang="ts">
    import Icon from "@iconify/svelte";

    type RatingProps = {
        fraction?: number,
        count?: number,
        value: number,
    }

    const {count = 5, value}: RatingProps = $props();
    const isFraction = value % 1 !== 0 && value % 1 >= 0.5;
</script>

<div class="flex gap-x-1 relative text-base">
    {#each Array.from({length: count}) as _, index}
        {#if isFraction && index === Math.floor(value)}
            <div class="relative">
                <div class="absolute left-0 top-0 w-[8px] overflow-hidden">
                    <Icon class="text-yellow-500 " icon="mingcute:star-fill"/>
                </div>
                <Icon class="text-neutral-500" icon="mingcute:star-fill"/>
            </div>
        {:else}
            <Icon class={index + 1 <= value ? 'text-yellow-500' : 'text-neutral-500'} icon="mingcute:star-fill"/>
        {/if}
    {/each}
</div>
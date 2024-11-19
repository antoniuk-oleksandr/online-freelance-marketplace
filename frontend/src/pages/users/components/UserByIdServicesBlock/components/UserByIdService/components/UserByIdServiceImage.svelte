<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";

    let {image, title}: UserService = $props();
    if (!image) image = '/images/no-image.jpg';
</script>

<img
        src={image}
        alt={title}
        class="object-cover aspect-video !rounded-md mb-3"
/>
<script lang="ts">
    import {getFileServerHost} from "@/utils/utils";

    type OrderSubmitRequirementsInfoImageProps = {
        title: string,
        image: string,
    }

    let {title, image} :OrderSubmitRequirementsInfoImageProps = $props();
    image = `${getFileServerHost()}/${image}`;
</script>

<img
        alt={title}
        class="aspect-video rounded-xl md:rounded-sm w-full object-cover object-center"
        src={image}
/>

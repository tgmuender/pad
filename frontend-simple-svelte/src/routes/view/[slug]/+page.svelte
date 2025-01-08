<script>
    export let data;

</script>

<div class="container">
    <h1>{data.name}</h1>
    <small>{data.id}</small>


    <div class="card" style="width: 18rem;">

        <div class="card-body">
            <h5 class="card-title">Meals</h5>
            <p class="card-text">View and create detailed meal plans for {data.name}.</p>
            <a href="{data.name}/meals?id={data.id}" class="btn btn-primary">Meals</a>
        </div>
    </div>

</div>

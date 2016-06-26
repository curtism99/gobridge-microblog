$.ajaxSetup({
    dataType: "json",
    contentType: 'application/json;charset=UTF-8'
})

$(function() {
    $("form").submit(function(e){
        e.preventDefault();

        var data = {
            body: $("textarea").val()
        };

        $.ajax({
            type: 'POST',
            url: '/posts',
            data: JSON.stringify(data),
        })
            .then(function() {
                console.log("saved post");
                alert("Saved post!")
            })
            .fail(function(resp) {
                alert("Failed to add Post.");
            });
    });
});
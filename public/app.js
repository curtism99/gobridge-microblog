$.ajaxSetup({
    dataType: 'json',
    contentType: 'application/json;charset=UTF-8'
});

function updateList() {
    $("#post-list").empty();
    $.ajax({
            type: 'GET',
            url: '/posts'
        })
            .then(function(data) {
                data.forEach(function(p, i) {
                    var html = "<div class='post bg-success'><h2 style='font-style: italic;'>" + '"' + p.title + '"' + "</h2><p>" + p.body + "</p><br/><small class='time'>" + p.time + "</small></div><br/>";
                    html += "<button style='margin-top: 25px;' class='del btn btn-lg btn-danger' data-id='" + i + "'>Delete</button>";
                    $("#post-list").append(html);
                });
                $("#title-field").val('');
                $("textarea").val('');
            })
            .fail(function() {
                alert("Could not get posts!");
            });
}

$(function() {
    $("form").submit(function(e){
        e.preventDefault();

        var data = {
            body: $("textarea").val(),
            title: $("#title-field").val()
        };

        $.ajax({
            type: 'POST',
            url: '/posts',
            data: JSON.stringify(data),
        })
            .then(function() {
                updateList();
            })
            .fail(function() {
                alert("Failed to add Post.");
            });
    });

    $("#post-list").click(function(e){
        var target = $(e.originalEvent.target);
        if (!target.hasClass("del")) {
            return;
        }

        var id = $(target).data("id");
        console.log("Deleting post", id);

        $.ajax({
            type: 'DELETE',
            url: '/posts/'+id
        })
            .then(function() {
                updateList();
            })
            .fail(function() {
                alert("Failed to delete Post.");
            });
    });

    // Make the initial call to update our list of posts
    updateList();
});
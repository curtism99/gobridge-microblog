$.ajaxSetup({
    dataType: "json",
    contentType: 'application/json;charset=UTF-8'
})

function updateList() {
    $("#post-list").empty();
    $.ajax({
            type: 'GET',
            url: '/posts'
        })
            .then(function(data) {
                data.forEach(function(p) {
                    var html = "<div class='post'><h2>" + p.title + "</h2><p>" + p.body + "</p><small class='time'>" + p.time + "</small></div><br/><br/>";
                    $("#post-list").append(html);
                });
                $("input").val('');
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
            title: $("input").val()
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
});

updateList();
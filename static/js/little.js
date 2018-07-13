(function() {
    $("#contact").submit(function(e){
	e.preventDefault();

	console.log($("#name").val())

	$.ajax({
            type: "POST",
            url: "/",
	    dataType: 'json',
	    contentType: "application/json; charset=utf-8",
	    data : JSON.stringify({
//		person : {
		    "name": $("#name").val(),
		    "email": $("#email").val(),
		    "phone_number": $("#phone_number").val()
//		}
	    }) // The data passed to / by POST method
        }).done(function(response_data) {
	    console.log(response_data)
            // Render your result
	    $("#contact").hide();
	    var msg = "Form submitted successfully. Here is the data that was sent: </br>" +
		response_data["name"] + "</br>" +
		response_data["email"] + "</br>" +
		response_data["phone_number"] + "</br>";

            $("#result").html(msg);
        });
    });
})();

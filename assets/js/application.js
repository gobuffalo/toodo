require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap/dist/js/bootstrap.js");

$(() => {

  $(".completed-box").on("click", (e) => {
    $(e.target).closest("form").submit();
  });

});

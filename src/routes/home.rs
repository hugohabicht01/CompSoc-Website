use rocket::{*, http::ContentType};
use rocket_dyn_templates::{Template, context};

use super::RouteProvider;
use super::super::engines::rsx::*;


#[get("/")]
fn index() -> Template {
    Template::render("index", context!{title: "LU CSS"})
}


#[component_route(get, "/jsx", JSX)]
fn jsx() -> Html {
    let component = html! { <div id="component">{"Some component"}</div> };
    
    let text_var = "You can interpolate text variables";
    
    html! {
       <div>
          {"You can type text right into the elements"}
          { component }
          { text_var }
       </div>
    }
}

pub struct HomeRoutes;

impl RouteProvider for HomeRoutes {
    fn base_url() -> &'static str {
        "/"
    }

    fn routes() -> Vec<Route> {
        routes![index, jsx]
    }
}

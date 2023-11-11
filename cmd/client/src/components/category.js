import React from "react";
export default function Category(props) {
    return (
    <div className="categories">
         <h2 id="list-heading">{props.name}</h2>
            {props.subcategories.map((subCat, idx) =>
                    <div key={subCat.SubCategoryName.idx} className="btn-group">
                        <button
                        type="button"
                        key={subCat.SubCategoryName.idx}
                        className="btn btn__primary"
                        onClick={() => props.startQuiz(props.name, subCat.SubCategoryName, subCat.URLPrefix)}
                        >
                        {subCat.SubCategoryName}
                        </button>
                    </div>
            )
            }
    </div>
    );
  }
import {default as React, useId} from "react";

export default function AnswerCheck(props) {

  const id =  props.answer.text.length

    return (
        <div key={props.idx} className="form-check">
        <label className="form-check-label" key={`label`-props.idx}>
            {props.answer.text}
        </label>
        <input
          className="form-check-input"
          id={id}
          key={id}
          type="checkbox"
          value={props.idx}
          onChange={
            props.handleChange
          }
        >
        </input>
    </div>

    );
  }
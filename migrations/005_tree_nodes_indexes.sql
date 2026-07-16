CREATE INDEX tree_nodes_user_id_idx
ON tree_nodes(user_id);

CREATE INDEX tree_nodes_parent_id_idx
ON tree_nodes(parent_id);

CREATE INDEX tree_nodes_user_parent_idx
ON tree_nodes(user_id, parent_id);
